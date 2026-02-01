export type WeekDay =
  | "Monday"
  | "Tuesday"
  | "Wednesday"
  | "Thursday"
  | "Friday";
// | "Saturday"
// | "Sunday";

export type ScheduleEvent = {
  id: string;
  subject: string;
  start_time: string;
  end_time: string;
  day: WeekDay;
};
