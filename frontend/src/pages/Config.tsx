import About from "../components/Config/About"
import Basic from "../components/Config/Basic"
import Influx from "../components/Config/Influx"
import Prometheus from "../components/Config/Prometheus"
import Scan from "../components/Config/Scan"

function Config() {

  return (
    <div class="row">
      <div class="col-md">
        <Basic></Basic>
        <div class="mt-4">
          <Scan></Scan>
        </div>
      </div>
      <div class="col-md">
        <Influx></Influx>
        <div class="mt-4">
          <Prometheus></Prometheus>
        </div>
        <div class="mt-4">
          <About></About>
        </div>
      </div>
    </div>
  )
}

export default Config