import MacHistory from "../MacHistory"

function HistCard(_props: any) {

  return (
    <div class="card border-primary">
      <div class="card-header">History</div>
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