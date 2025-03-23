import { For } from "solid-js"
import Filter from "../components/Filter"
import { allHosts, setShow, show } from "../functions/exports"
import MacHistory from "../components/MacHistory"
import HistShow from "../components/HistShow"

function History() {

  const showStr = localStorage.getItem("histShow") as string;
  setShow(+showStr);
  (show() === 0 || isNaN(show())) ? setShow(198) : '';
  
  return (
    <div class="card border-primary">
      <div class="card-header d-flex justify-content-between">
        <Filter></Filter>
        <HistShow name="histShow"></HistShow>
      </div>
      <div class="card-body">
        <table class="table table-striped table-hover">
          <tbody>
          <For each={allHosts()}>{(host, index) =>
          <tr>
            <td class="opacity-50" style="width: 2em;">{index()}.</td>
            <td>
              <a href={"/host/"+host.ID}>{host.Name}</a><br></br>
              <a href={"http://"+host.IP}>{host.IP}</a>
            </td>
            <td>
              <MacHistory mac={host.Mac}></MacHistory>
            </td>
          </tr>
          }</For>
          </tbody> 
        </table>
      </div>
    </div>
  )
}

export default History
