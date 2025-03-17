import { allHosts } from "../functions/exports";
import { For } from "solid-js";
import TableRow from "../components/TableRow";
import CardHead from "../components/CardHead";
import TableHead from "../components/TableHead";


function Body() {

  return (
    <div class="card border-primary">
      <div class="card-header">
        <CardHead></CardHead>
      </div>
      <div class="card-body table-responsive">
        <table class="table table-striped table-hover">
          <TableHead></TableHead>
          <tbody>
            <For each={allHosts()}>{(host, index) =>
            <TableRow host={host} index={index}></TableRow>
            }</For>
          </tbody> 
        </table>
      </div>
    </div>
  )
}

export default Body
