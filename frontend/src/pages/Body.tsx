import { For } from "solid-js";

import { allHosts } from "../functions/exports";

import TableRow from "../components/Body/TableRow";
import TableHead from "../components/Body/TableHead";
import CardHead from "../components/Body/CardHead";

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
            <TableRow host={host} index={index() + 1}></TableRow>
            }</For>
          </tbody> 
        </table>
      </div>
    </div>
  )
}

export default Body
