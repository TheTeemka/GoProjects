"use client";

import React, { useState } from "react";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { toast } from "@/lib/toast";
import type {
  CreateScheduleRequest,
  ScheduleEvent,
  WeekDay,
} from "@/types/schedule";
import { Plus } from "lucide-react";

export interface AddScheduleModalProps {
  trigger?: React.ReactNode;
  onConfirm?: (req: Omit<ScheduleEvent, "id">) => Promise<boolean>;
  onSuccess?: () => void;
}

const weekDays: WeekDay[] = [
  "Monday",
  "Tuesday",
  "Wednesday",
  "Thursday",
  "Friday",
];

export function AddScheduleModal({
  trigger,
  onConfirm,
  onSuccess,
}: AddScheduleModalProps) {
  const [open, setOpen] = useState(false);
  const [subject, setSubject] = useState("");
  const [day, setDay] = useState<number>(-1);
  const [startTime, setStartTime] = useState("00:00");
  const [endTime, setEndTime] = useState("00:00");
  const [loading, setLoading] = useState(false);

  const reset = () => {
    setSubject("");
    setDay(-1);
    setStartTime("00:00");
    setEndTime("00:00");
  };

  const handleSubmit = async () => {
    console.log({ subject, day, startTime, endTime });
    if (!subject || day === -1 || !startTime || !endTime) {
      toast.warn("Please fill all fields");
      return;
    }

    if (!onConfirm) {
      toast.error("Create handler not provided");
      return;
    }

    setLoading(true);
    try {
      const req: CreateScheduleRequest = {
        subject,
        day_of_week: day,
        start_time: startTime,
        end_time: endTime,
        group_id: 1,
      };
      const success = await onConfirm(req);
      if (success) {
        toast.success("Schedule created");
        setOpen(false);
        reset();
        onSuccess?.();
      } else {
        toast.error("Failed to create schedule");
      }
    } catch (err) {
      toast.error(
        err instanceof Error ? err.message : "Failed to create schedule",
      );
    } finally {
      setLoading(false);
    }
  };

  const triggerNode = (trigger as React.ReactNode) ?? (
    <Button className="ml-auto">
      <Plus className="mr-2 h-4 w-4" />
      Add Event
    </Button>
  );

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{triggerNode}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add Schedule Event</DialogTitle>
          <DialogDescription>
            Add a new schedule event to the group calendar.
          </DialogDescription>
        </DialogHeader>

        <div className="p-4 space-y-4">
          <div className="grid gap-2">
            <Label>Subject</Label>
            <Input
              value={subject}
              onChange={(e) => setSubject(e.target.value)}
            />
          </div>

          <div className="grid gap-2">
            <Label>Day</Label>
            <select
              value={day}
              onChange={(e) => setDay(Number(e.target.value))}
              className="h-9 w-full rounded-md border bg-transparent px-3 text-base"
            >
              <option value={-1}>Select day</option>
              {weekDays.map((d, index) => (
                <option key={d} value={index}>
                  {d}
                </option>
              ))}
            </select>
          </div>

          <div className="grid gap-2">
            <Label>Start time</Label>
            <Input
              type="time"
              value={startTime}
              onChange={(e) => setStartTime(e.target.value)}
            />
          </div>

          <div className="grid gap-2">
            <Label>End time</Label>
            <Input
              type="time"
              value={endTime}
              onChange={(e) => setEndTime(e.target.value)}
            />
          </div>
        </div>

        <DialogFooter>
          <div className="flex gap-2 w-full">
            <Button variant="outline" onClick={() => setOpen(false)}>
              Cancel
            </Button>
            <Button
              onClick={handleSubmit}
              className="ml-auto"
              disabled={loading}
            >
              {loading ? "Saving..." : "Save"}
            </Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export default AddScheduleModal;
