import { For, onCleanup, onMount, Show } from "solid-js";
import { getHistoryForMac } from "../functions/history";
import { Host, show } from "../functions/exports";
import { createStore } from "solid-js/store";

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

  return (
    <For each={hist}>{(h, index) =>
      <Show
        when={index() < show()}
      >
        <i title={"Date:"+h.Date+"\nIface:"+h.Iface+"\nIP:"+h.IP+"\nKnown:"+h.Known} 
        class={h.Now === 0?"my-box-off":"my-box-on"}></i>
      </Show>
    }</For>
  )
}

export default MacHistory
