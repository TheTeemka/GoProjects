export type WeekDay =
  | "Monday"
  | "Tuesday"
  | "Wednesday"
  | "Thursday"
  | "Friday";

export type ScheduleEvent = {
  id: string;
  group_id: number;
  subject: string;
  start_time: string;
  end_time: string;
  day_of_week: number;
};

export type CreateScheduleRequest = Omit<ScheduleEvent, "id">;
