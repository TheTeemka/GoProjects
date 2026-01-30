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
  startTime: string;
  endTime: string;
  day: WeekDay;
};
