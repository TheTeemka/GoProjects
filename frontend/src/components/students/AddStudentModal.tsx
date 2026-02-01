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
import { Plus } from "lucide-react";
import type { CreateStudentRequest } from "@/types/student";

export interface AddStudentModalProps {
  defaultGroupId?: number;
  trigger?: React.ReactNode;
  onConfirm?: (req: CreateStudentRequest) => Promise<boolean>;
  onSuccess?: () => void;
}

export function AddStudentModal({
  defaultGroupId = 1,
  trigger,
  onConfirm,
  onSuccess,
}: AddStudentModalProps) {
  const [open, setOpen] = useState(false);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [birthday, setBirthday] = useState("");
  const [group_id, setGroupId] = useState<number>(defaultGroupId);
  const [loading, setLoading] = useState(false);

  const reset = () => {
    setName("");
    setEmail("");
    setBirthday("");
    setGroupId(defaultGroupId);
  };

  const handleSubmit = async () => {
    if (!name || !email || !birthday || !group_id) {
      toast.warn("Please fill all fields");
      return;
    }

    if (!onConfirm) {
      toast.error("Create handler not provided");
      return;
    }

    setLoading(true);
    try {
      const formattedBirthday = new Date(birthday).toISOString();
      const req: CreateStudentRequest = {
        name,
        email,
        birthday: formattedBirthday,
        group_id,
      };
      const success = await onConfirm(req);
      if (success) {
        toast.success("Student created");
        setOpen(false);
        reset();
        onSuccess?.();
      } else {
        toast.error("Failed to create student");
      }
    } catch (err) {
      toast.error(
        err instanceof Error ? err.message : "Failed to create student",
      );
    } finally {
      setLoading(false);
    }
  };

  const triggerNode = (trigger as React.ReactNode) ?? (
    <Button className="ml-auto">
      <Plus className="mr-2 h-4 w-4" />
      Add Student
    </Button>
  );

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>{triggerNode}</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Add Student</DialogTitle>
          <DialogDescription>
            Fill in student details and save.
          </DialogDescription>
        </DialogHeader>

        <div className="p-4 space-y-4">
          <div className="grid gap-2">
            <Label>Name</Label>
            <Input value={name} onChange={(e) => setName(e.target.value)} />
          </div>

          <div className="grid gap-2">
            <Label>Email</Label>
            <Input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div className="grid gap-2">
            <Label>Birthday</Label>
            <Input
              type="date"
              value={birthday}
              onChange={(e) => setBirthday(e.target.value)}
            />
          </div>

          <div className="grid gap-2">
            <Label>Group ID</Label>
            <Input
              type="number"
              value={String(group_id ?? "")}
              onChange={(e) => {
                if (e.target.value !== "") {
                  setGroupId(Number(e.target.value));
                } else {
                  setGroupId(0);
                }
              }}
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

export default AddStudentModal;
