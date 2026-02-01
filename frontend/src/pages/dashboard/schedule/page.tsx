import { useState } from "react";
import { ScheduleCalendar } from "@/components/schedule/ScheduleCalendar";
import AddScheduleModal from "@/components/schedule/AddScheduleModal";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { useSchedules } from "@/hooks/useSchedules";
import DeleteScheduleModal from "@/components/schedule/DeleteScheduleModal";

const SchedulePage = () => {
  const { events, createSchedule, deleteSchedule } = useSchedules(1);
  const [deletingId, setDeletingId] = useState<string | null>(null);
  return (
    <div className="space-y-6 px-10 py-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Schedule</h1>
          <p className="text-gray-500 mt-2">Manage group schedule</p>
        </div>
        <div>
          <AddScheduleModal
            trigger={
              <Button>
                <Plus className="mr-2 h-4 w-4" />
                Add Event
              </Button>
            }
            onConfirm={createSchedule}
          />
        </div>
      </div>

      <ScheduleCalendar
        events={events}
        onDeleteEventClick={(e) => setDeletingId(e.id)}
      />

      <DeleteScheduleModal
        id={deletingId}
        onConfirm={deleteSchedule}
        onSuccess={() => setDeletingId(null)}
        onClose={() => setDeletingId(null)}
      />
    </div>
  );
};

export default SchedulePage;
