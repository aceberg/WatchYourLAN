import { For, onCleanup, onMount, Show } from "solid-js";
import { getHistoryForMac } from "../functions/history";
import { Host, show } from "../functions/exports";
import { createStore } from "solid-js/store";
import { getHistoryPeriod, historyPeriodLabel, parseHistoryTimestamp } from "../functions/historyPeriods";

function MacHistory(_props: any) {

  const [hist, setHist] = createStore<Host[]>([]);
  let interval: number;

  onMount(async () => {
    const newHistory = await getHistoryForMac(_props.mac, _props.date);
    setHist(newHistory);
    interval = setInterval(async () => {
      // console.log("Upd Hist", new Date());
      const newHistory = await getHistoryForMac(_props.mac, _props.date);
      setHist(newHistory);
    }, 60000); // 60000 ms = 1 minute
  });

  onCleanup(() => {
    clearInterval(interval);
  });

  const statusLabel = (host: Host) => host.Now === 0 ? "Offline" : "Online";
  const knownLabel = (host: Host) => host.Known === 0 ? "Unknown" : "Known";

  const sampleTitle = (host: Host) => {
    const period = host.Now === 1 ? historyPeriodLabel(getHistoryPeriod(host.Date)) : "";

    return "Date: " + host.Date
      + "\nStatus: " + statusLabel(host)
      + (period ? "\nPeriod: " + period : "")
      + "\nIface: " + host.Iface
      + "\nIP: " + host.IP
      + "\nKnown: " + knownLabel(host);
  };

  const sampleClass = (host: Host) => {
    if (host.Now === 0) {
      return "my-box-off";
    }

    const period = getHistoryPeriod(host.Date);

    if (period === "night") {
      return "my-box-on my-box-on-night";
    }

    return period === "day" ? "my-box-on my-box-on-day" : "my-box-on";
  };

  const boundaryClass = (host: Host, index: number) => {
    if (index === 0) {
      return "";
    }

    const previous = hist[index - 1];

    if (!previous) {
      return "";
    }

    const currentDate = parseHistoryTimestamp(host.Date);
    const previousDate = parseHistoryTimestamp(previous.Date);

    if (!currentDate || !previousDate) {
      return "";
    }

    if (currentDate.year !== previousDate.year
      || currentDate.month !== previousDate.month
      || currentDate.day !== previousDate.day) {
      return " history-sample-day-break";
    }

    if (currentDate.hour !== previousDate.hour) {
      return " history-sample-hour-break";
    }

    return "";
  };

  return (
    <For each={hist}>{(h, index) =>
      <Show
        when={index() < show()}
      >
        <i
          title={sampleTitle(h)}
          aria-label={sampleTitle(h)}
          role="img"
          class={sampleClass(h) + " history-sample" + boundaryClass(h, index())}
        ></i>
      </Show>
    }</For>
  )
}

export default MacHistory
