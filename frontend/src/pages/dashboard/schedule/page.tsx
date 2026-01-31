import { useState } from "react";
import { ScheduleCalendar } from "@/components/schedule/ScheduleCalendar";
import type { ScheduleEvent } from "@/types/schedule";

const SchedulePage = () => {
  const [events] = useState<ScheduleEvent[]>([
    {
      id: "1",
      subject: "Team Meeting",
      start_time: "09:00",
      end_time: "10:00",
      day: "Monday",
    },
    {
      id: "2",
      subject: "Project Review",
      start_time: "14:00",
      end_time: "15:30",
      day: "Monday",
    },
    {
      id: "3",
      subject: "Client Meeting",
      start_time: "11:00",
      end_time: "12:00",
      day: "Wednesday",
    },
    {
      id: "4",
      subject: "Development Sprint",
      start_time: "10:00",
      end_time: "12:00",
      day: "Friday",
    },
  ]);

  return <ScheduleCalendar events={events} />;
};

export default SchedulePage;
