import type { ScheduleEvent } from "@/types/schedule";
import apiClient from "./axiosClient";

export type CreateScheduleRequest = Omit<ScheduleEvent, "id">;
export type UpdateScheduleRequest = Partial<CreateScheduleRequest>;

const getSchedulesByID = async (groupId: number): Promise<ScheduleEvent[]> => {
  const resp = await apiClient.get<ScheduleEvent[]>(
    `/schedules/group/${groupId}`,
  );
  return resp.data;
};

const createSchedule = async (req: CreateScheduleRequest): Promise<boolean> => {
  const resp = await apiClient.post("/schedules", req);
  return resp.status === 201;
};

const updateSchedule = async (
  id: string,
  req: UpdateScheduleRequest,
): Promise<boolean> => {
  const resp = await apiClient.put(`/schedules/${id}`, req);
  return resp.status === 200;
};

const deleteSchedule = async (id: string): Promise<boolean> => {
  const resp = await apiClient.delete(`/schedules/${id}`);
  return resp.status === 200 || resp.status === 204;
};

export const schedulesApi = {
  getSchedulesByID,
  createSchedule,
  updateSchedule,
  deleteSchedule,
};

export default schedulesApi;
