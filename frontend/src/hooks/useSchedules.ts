import schedulesApi from "@/api/schedules";
import type { ScheduleEvent } from "@/types/schedule";
import { useCallback, useEffect, useState } from "react";

export function useSchedules(initialGroupID: number) {
  const [groupID, setGroupID] = useState<number>(initialGroupID);
  const [events, setEvents] = useState<ScheduleEvent[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const fetchSchedules = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const data = await schedulesApi.getSchedulesByID(groupID);
      setEvents(data);
      return data;
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Failed to fetch schedules",
      );
      return [] as ScheduleEvent[];
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchSchedules();
  }, [groupID]);

  const createSchedule = async (
    event: Omit<ScheduleEvent, "id">,
  ): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const success = await schedulesApi.createSchedule(event);
      if (success) {
        await fetchSchedules();
        return true;
      } else {
        setError("Failed to create schedule");
        return false;
      }
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Failed to create schedule",
      );
      return false;
    } finally {
      setLoading(false);
    }
  };

  const deleteSchedule = async (id: string): Promise<boolean> => {
    setLoading(true);
    setError(null);
    try {
      const success = await schedulesApi.deleteSchedule(id);
      if (success) {
        await fetchSchedules();
        return true;
      } else {
        setError("Failed to delete schedule");
        return false;
      }
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Failed to delete schedule",
      );
      return false;
    } finally {
      setLoading(false);
    }
  };

  return {
    events,
    loading,
    error,
    groupID,
    setGroupId: setGroupID,
    fetchSchedules,
    createSchedule,
    deleteSchedule,
  };
}
