import { createEffect, For, Show } from "solid-js"
import Filter from "../components/Filter"
import { allHosts, histUpdOnFilter, Host, setHistUpdOnFilter, setShow, show } from "../functions/exports"
import MacHistory from "../components/MacHistory"
import HistShow from "../components/HistShow"

function History() {

  let hosts: Host[] = [];
  hosts.push(...allHosts);

  const showStr = localStorage.getItem("histShow") as string;
  setShow(+showStr);
  (show() === 0 || isNaN(show())) ? setShow(200) : '';
  
  createEffect(() => {
    if (histUpdOnFilter()) {
      hosts = [];
      hosts.push(...allHosts);
      console.log("Upd on Filter");
      setHistUpdOnFilter(false);
    }
  });

  return (
    <div class="card border-primary">
      <div class="card-header d-flex justify-content-between">
        <Filter></Filter>
        <HistShow name="histShow"></HistShow>
      </div>
      <div class="card-body">
        <table class="table table-striped table-hover">
          <tbody>
          <Show
            when={!histUpdOnFilter()}
          >
            <For each={hosts}>{(host, index) =>
            <tr>
              <td class="opacity-50" style="width: 2em;">{index()+1}.</td>
              <td>
                <a href={"/host/"+host.ID}>{host.Name}</a><br></br>
                <a href={"http://"+host.IP}>{host.IP}</a>
              </td>
              <td>
                <MacHistory mac={host.Mac}></MacHistory>
              </td>
            </tr>
            }</For>
          </Show>
          </tbody> 
        </table>
      </div>
    </div>
  )
}

export default History
