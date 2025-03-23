import { For } from "solid-js"
import Filter from "../components/Filter"
import { allHosts } from "../functions/exports"
import MacHistory from "../components/MacHistory"

function History() {

  return (
    <div class="card border-primary">
      <div class="card-header d-flex justify-left">
        <Filter></Filter>
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
              <MacHistory mac={host.Mac} show={200}></MacHistory>
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
