import { Show } from "solid-js";
import { editNames, setEditNames } from "../../functions/exports";
import Filter from "../Filter";
import Search from "../Search";

function CardHead() {

  return (
    <div class="row">
      <div class="col-md mt-1 mb-1">
        <Filter></Filter>
      </div>
      <div class="col-md mt-1 mb-1">
        <div class="d-flex justify-content-between">
        <Search></Search>
        <Show
          when={editNames()}
          fallback={<button class="btn btn-outline-primary" onClick={[setEditNames, true]}>Edit names</button>}
        >
          <button class="btn btn-primary" onClick={[setEditNames, false]}>Edit names</button>
        </Show>
        </div>
      </div>
    </div>
  )
}

export default CardHead
