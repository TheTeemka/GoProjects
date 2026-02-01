"use client";

import { useEffect, useState } from "react";
import {
  Dialog,
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
import type { Student, UpdateStudentRequest } from "@/types/student";

export interface UpdateStudentModalProps {
  student?: Student | null;
  onConfirm?: (id: number, req: UpdateStudentRequest) => Promise<boolean>;
  onSuccess?: () => void;
  onClose?: () => void;
}

export function UpdateStudentModal({
  student,
  onConfirm,
  onSuccess,
  onClose,
}: UpdateStudentModalProps) {
  const [open, setOpen] = useState(false);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [birthday, setBirthday] = useState("");
  const [group_id, setGroupId] = useState<number>(0);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (student) {
      setName(student.name);
      setEmail(student.email);
      setBirthday(student.birthday);
      setGroupId(student.group_id);
      setOpen(true);
    } else {
      setOpen(false);
    }
  }, [student]);

  const handleSubmit = async () => {
    if (!student) return;
    if (!name || !email || !birthday || !group_id) {
      toast.warn("Please fill all fields");
      return;
    }

    if (!onConfirm) {
      toast.error("Update handler not provided");
      return;
    }

    setLoading(true);
    try {
      const req: UpdateStudentRequest = {
        name,
        email,
        birthday,
        group_id,
      };
      const success = await onConfirm(student.id, req);
      if (success) {
        toast.success("Student updated");
        onSuccess?.();
        setOpen(false);
        onClose?.();
      } else {
        toast.error("Failed to update student");
      }
    } catch (err) {
      toast.error(
        err instanceof Error ? err.message : "Failed to update student",
      );
    } finally {
      setLoading(false);
    }
  };

  return (
    <Dialog
      open={open}
      onOpenChange={(v) => {
        setOpen(v);
        if (!v) onClose?.();
      }}
    >
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Edit Student</DialogTitle>
          <DialogDescription>
            Update student details and save.
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
              type="text"
              placeholder="YYYY-MM-DD"
              value={birthday}
              onChange={(e) => setBirthday(e.target.value)}
            />
          </div>

          <div className="grid gap-2">
            <Label>Group ID</Label>
            <Input
              type="number"
              value={String(group_id)}
              onChange={(e) => setGroupId(Number(e.target.value))}
            />
          </div>
        </div>

        <DialogFooter>
          <div className="flex gap-2 w-full">
            <Button
              variant="outline"
              onClick={() => {
                setOpen(false);
                onClose?.();
              }}
            >
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

export default UpdateStudentModal;
