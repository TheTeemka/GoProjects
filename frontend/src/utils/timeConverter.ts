export function HoursAndMinutesToDecimalTime(time: string): number {
  const [hoursStr, minutesStr] = time.split(":");
  const hours = parseInt(hoursStr, 10);
  const minutes = parseInt(minutesStr, 10);
  return hours + minutes / 60;
}
