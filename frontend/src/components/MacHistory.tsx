import { createSignal, For, onMount, Show } from "solid-js";
import { apiGetHistory } from "../functions/api";
import { Host, show } from "../functions/exports";

function MacHistory(_props: any) {

  const [hist, setHist] = createSignal<Host[]>([]);

  onMount(async () => {
    let h:Host[] = [];
    h = await apiGetHistory(_props.mac);
    if (h != null) {
      h.sort((a:Host, b:Host) => (a.Date < b.Date ? 1 : -1));
      setHist(h);
    }
  });

  return (
    <For each={hist()}>{(h, index) =>
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
