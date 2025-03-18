import { For } from "solid-js";

import { allHosts } from "../functions/exports";

import TableRow from "../components/TableRow";
import TableHead from "../components/TableHead";
import BodyCardHead from "../components/BodyCardHead";

function Body() {

  return (
    <div class="card border-primary">
      <div class="card-header">
        <BodyCardHead></BodyCardHead>
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
