import { useState } from "react";
import { Button } from "@/components/ui/button";
import { ScheduleCalendar } from "@/components/schedule/ScheduleCalendar";
import type { ScheduleEvent } from "@/types/schedule";

const SchedulePage = () => {
  const [events] = useState<ScheduleEvent[]>([
    {
      id: "1",
      subject: "Team Meeting",
      startTime: "09:00",
      endTime: "10:00",
      day: "Monday",
    },
    {
      id: "2",
      subject: "Project Review",
      startTime: "14:00",
      endTime: "15:30",
      day: "Monday",
    },
    {
      id: "3",
      subject: "Client Meeting",
      startTime: "11:00",
      endTime: "12:00",
      day: "Wednesday",
    },
    {
      id: "4",
      subject: "Development Sprint",
      startTime: "10:00",
      endTime: "12:00",
      day: "Friday",
    },
  ]);

  return <ScheduleCalendar events={events} />;
};

export default SchedulePage;
