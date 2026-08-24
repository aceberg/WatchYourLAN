export function formatLastSeen(date: string) {
  const match = /^(\d{4})-(\d{2})-(\d{2})[ T](\d{2}):(\d{2})(?::\d{2})?$/.exec(date);

  if (!match) {
    return date;
  }

  const [, year, month, day, hour, minute] = match;
  const monthName = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"][Number(month) - 1];
  const currentYear = new Date().getFullYear().toString();

  if (!monthName) {
    return date;
  }

  if (year === currentYear) {
    return `${day} ${monthName} ${hour}:${minute}`;
  }

  return `${day} ${monthName} ${year} ${hour}:${minute}`;
}
