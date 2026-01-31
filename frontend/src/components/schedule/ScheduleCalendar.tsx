import type { ScheduleEvent, WeekDay } from "@/types/schedule";
import React from "react";
import { HoursAndMinutesToDecimalTime } from "@/lib/timeConverter";
// ...existing code...
const SUBJECT_COLOR_PALETTE: string[] = [
  "bg-red-400",
  "bg-rose-400",
  "bg-pink-400",
  "bg-fuchsia-400",
  "bg-purple-400",
  "bg-indigo-400",
  "bg-sky-400",
  "bg-teal-400",
  "bg-emerald-400",
  "bg-amber-400",
];

function getColorForSubject(subject: string): string {
  let hash = 5381;
  for (let i = 0; i < subject.length; i++) {
    hash = (hash * 33) ^ subject.charCodeAt(i);
  }
  const idx = Math.abs(hash) % SUBJECT_COLOR_PALETTE.length;
  console.log(idx);
  return SUBJECT_COLOR_PALETTE[idx];
}

export function ScheduleCalendar({ events }: { events: ScheduleEvent[] }) {
  const eventsByDay: Record<WeekDay, ScheduleEvent[]> = {
    Monday: [],
    Tuesday: [],
    Wednesday: [],
    Thursday: [],
    Friday: [],
    // Saturday: [],
    // Sunday: [],
  };

  for (const event of events) {
    eventsByDay[event.day].push(event);
  }

  const daysOfWeek: WeekDay[] = [
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    // "Saturday",
    // "Sunday",
  ];

  const calendarStartHour = 8;
  const hours = Array.from(
    { length: 24 - calendarStartHour },
    (_, i) => i + calendarStartHour,
  );

  // horizontal lines (rows) + vertical lines between day columns
  const rightBackground = {
    backgroundImage: `
      repeating-linear-gradient(to bottom, rgba(0,0,0,0.06) 0px 1px, transparent 1px var(--height-event-card)),
      repeating-linear-gradient(to right, rgba(0,0,0,0.06) 0px 1px, transparent 1px calc(100% / ${daysOfWeek.length}))
    `,
    // offset both patterns by header height so lines start below headers
    backgroundPosition: `0px var(--height-header), 0px var(--height-header)`,
    backgroundRepeat: "repeat, repeat",
  } as React.CSSProperties;

  // hours column should also show horizontal lines using the same var
  const hoursColumnStyle: React.CSSProperties = {
    backgroundImage: `repeating-linear-gradient(to bottom, rgba(0,0,0,0.06) 0px 1px, transparent 1px var(--height-event-card))`,
    backgroundPosition: `0px var(--height-header)`,
    backgroundRepeat: "repeat",
  };

  return (
    <div className="min-h-screen dark:bg-gray-900">
      <div className="max-w-7xl mx-auto">
        <div className="bg-white flex dark:bg-gray-800 rounded-lg shadow overflow-hidden">
          {/* hours column */}
          <div
            className="w-20 shrink-0 border-r border-gray-200 dark:border-gray-700"
            style={hoursColumnStyle}
          >
            {/* header spacer */}
            <div className="h-header" />
            {hours.map((hour) => (
              <div key={hour} className="flex items-start px-2 h-event-card">
                <div className="text-sm text-gray-500 dark:text-gray-400 text-center">
                  {hour}:00
                </div>
              </div>
            ))}
          </div>

          {/* days area (horizontal + vertical grid lines via background) */}
          <div className="flex-1 overflow-auto" style={rightBackground}>
            {/* day headers */}
            <div
              style={{
                display: "grid",
                gridTemplateColumns: `repeat(${daysOfWeek.length}, minmax(0, 1fr))`,
              }}
            >
              {daysOfWeek.map((day, index) => {
                const today = new Date();
                const isToday = index === (today.getDay() + 6) % 7;
                return (
                  <div>
                    <div
                      key={day}
                      className={`h-header p-2 text-center bg-gray-50 dark:bg-gray-900 ${
                        isToday ? "border-b-2 border-blue-500" : ""
                      }`}
                    >
                      <div className="font-semibold text-gray-900 dark:text-white">
                        {day}
                      </div>
                    </div>
                    <div key={day} className="relative">
                      {isToday && <CurrentTimeIndicator />}
                      {eventsByDay[day].map((event) => (
                        <EventCard
                          key={event.id}
                          event={event}
                          color={getColorForSubject(event.subject)}
                        />
                      ))}
                    </div>
                  </div>
                );
              })}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

interface EventCardProps {
  event: ScheduleEvent;
  color?: string;
}

function EventCard({ event, color = "bg-chart-1" }: EventCardProps) {
  const startHour = HoursAndMinutesToDecimalTime(event.start_time);
  const endHour = HoursAndMinutesToDecimalTime(event.end_time);
  const durationHours = endHour - startHour;

  console.log({ startHour, endHour, durationHours });
  const positionStyle: React.CSSProperties = {
    top: `calc(var(--height-event-card) * ${startHour - 8})`,
    height: `calc(var(--height-event-card) * ${durationHours})`,
  };
  return (
    <div
      className={`${color} relative  text-white text-xs p-2 rounded mb-1 cursor-pointer hover:opacity-90 transition-opacity`}
      style={positionStyle}
    >
      <div className="font-semibold">{event.subject}</div>
      <div className="text-xs opacity-90">
        {event.start_time} - {event.end_time}
      </div>
    </div>
  );
}

function CurrentTimeIndicator() {
  const today = new Date();
  const hours = today.getHours() - 8;
  const minutes = today.getMinutes();
  const decimalTime = hours + minutes / 60;

  const topPosition = `calc(var(--height-event-card) * ${decimalTime})`;

  return (
    <div
      className="absolute left-0 right-0 h-0.5 bg-red-500"
      style={{ top: topPosition }}
    />
  );
}
