import { createSignal } from "solid-js"
import { apiPath } from "../../functions/api"
import { saveErrorMessage, submitConfigForm } from "../../functions/configForms";
import { appConfig } from "../../functions/exports"

function Prometheus() {
  const [status, setStatus] = createSignal("");
  const [error, setError] = createSignal("");

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    const form = event.currentTarget as HTMLFormElement;
    setStatus("");
    setError("");

    try {
      await submitConfigForm(form);
      setStatus("Saved");
    } catch (saveError) {
      setError(saveErrorMessage(saveError));
    }
  };

  return (
    <div class="card wyl-panel config-panel">
      <div class="card-header">Prometheus config</div>
      <div class="card-body table-responsive">
        <form action={apiPath + '/api/config_prometheus/'} method="post" onSubmit={handleSubmit}>
          <table class="table table-borderless"><tbody>
            <tr>
              <td class="config-field-label">Enable</td>
              <td class="config-field-value">
                <div class="form-check form-switch">
                  {appConfig().PrometheusEnable
                    ? <input class="form-check-input" type="checkbox" name="enable" checked></input>
                    : <input class="form-check-input" type="checkbox" name="enable"></input>
                  }
                </div>
              </td>
            </tr>
            <tr>
              <td class="config-action-cell">
                <button type="submit" class="btn btn-sm wyl-button">Save Prometheus</button>
                <span class={"config-save-status" + (error() ? " config-save-error" : "")} role="status">
                  {error() || status()}
                </span>
              </td>
              <td class="config-action-cell">
                <a href="/metrics" target="_blank">/metrics</a>
              </td>
            </tr>
          </tbody></table>
        </form>
      </div>
    </div>
  )
}

export default Prometheus
