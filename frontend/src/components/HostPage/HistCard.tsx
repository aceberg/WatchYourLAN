import { setShow, show } from "../../functions/exports";
import HistShow from "../HistShow"
import MacHistory from "../MacHistory"

function HistCard(_props: any) {

  const showStr = localStorage.getItem("hostShow") as string;
  setShow(+showStr);
  (show() === 0 || isNaN(show())) ? setShow(500) : '';

  return (
    <div class="card border-primary">
      <div class="card-header d-flex justify-content-between">
        <div>Host History</div>
        <HistShow name="hostShow"></HistShow>
      </div>
      <div class="card-body">
      {_props.mac !== ""
      ? <MacHistory mac={_props.mac}></MacHistory>
      : <>Loading...</>
      }
      </div>
    </div>
  )
}

export default HistCard