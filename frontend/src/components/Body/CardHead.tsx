import { Show } from "solid-js";
import { editNames, setEditNames } from "../../functions/exports";
import Filter from "../Filter";
import Search from "../Search";
import { getHosts } from "../../functions/atstart";

function CardHead() {

  const handleEditNames = (toggle: boolean) => {
    if (!toggle) {
      getHosts();
    }
    setEditNames(toggle);
  };

  return (
    <div class="row">
      <div class="col-md mt-1 mb-1">
        <div class="d-flex justify-left">
        <Filter></Filter>
        </div>
      </div>
      <div class="col-md mt-1 mb-1">
        <div class="d-flex justify-content-between">
        <Search></Search>
        <Show
          when={editNames()}
          fallback={<button class="btn btn-outline-primary" title="Toggle edit" onClick={[handleEditNames, true]}>Edit</button>}
        >
          <button class="btn btn-primary" title="Toggle edit" onClick={[handleEditNames, false]}>Edit</button>
        </Show>
        </div>
      </div>
    </div>
  )
}

export default CardHead
