"use client";

import React, { useEffect, useState } from "react";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { toast } from "@/lib/toast";

export interface DeleteStudentModalProps {
  id?: number | null;
  onConfirm?: (id: number) => Promise<boolean>;
  onSuccess?: () => void;
  onClose?: () => void;
}

export function DeleteStudentModal({
  id,
  onConfirm,
  onSuccess,
  onClose,
}: DeleteStudentModalProps) {
  const [open, setOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    setOpen(typeof id !== "undefined" && id !== null);
  }, [id]);

  const handleDelete = async () => {
    if (!id) return;
    if (!onConfirm) {
      toast.error("Delete handler not provided");
      return;
    }
    setLoading(true);
    try {
      const success = await onConfirm(id);
      if (success) {
        toast.success("Student deleted");
        onSuccess?.();
        setOpen(false);
        onClose?.();
      } else {
        toast.error("Failed to delete student");
      }
    } catch (err) {
      toast.error(
        err instanceof Error ? err.message : "Failed to delete student",
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
          <DialogTitle>Delete Student</DialogTitle>
          <DialogDescription>
            Are you sure you want to delete this student?
          </DialogDescription>
        </DialogHeader>

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
              onClick={handleDelete}
              className="ml-auto"
              variant="destructive"
              disabled={loading}
            >
              {loading ? "Deleting..." : "Delete"}
            </Button>
          </div>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export default DeleteStudentModal;
