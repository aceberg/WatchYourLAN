import { createSignal } from "solid-js";
import { apiSetRetention } from "../../functions/api";
import { appConfig, setAppConfig } from "../../functions/exports";

function Retention() {
  const [status, setStatus] = createSignal("");
  const [error, setError] = createSignal("");

  const handleSubmit = async (event: SubmitEvent) => {
    event.preventDefault();

    const form = event.currentTarget as HTMLFormElement;
    const data = new FormData(form);
    const presenceRetention = Number(data.get("presenceRetention"));
    const connectivityRetention = Number(data.get("connectivityRetention"));

    setStatus("");
    setError("");

    try {
      const nextConfig = await apiSetRetention(presenceRetention, connectivityRetention);
      setAppConfig(nextConfig);
      setStatus("Saved");
    } catch {
      setError("Retention values must be positive whole hours.");
    }
  };

  return (
    <div class="card wyl-panel config-panel">
      <div class="card-header">Data retention</div>
      <div class="card-body table-responsive">
        <form onSubmit={handleSubmit}>
          <table class="table table-borderless">
            <tbody>
              <tr>
                <td class="config-field-label config-field-label-top">Presence retention</td>
                <td class="config-field-value">
                  <div class="config-value-with-unit">
                    <input
                      name="presenceRetention"
                      type="number"
                      min="1"
                      step="1"
                      class="form-control config-retention-input"
                      value={appConfig().TrimHist}
                      required
                    ></input>
                    <span class="config-field-unit">hours</span>
                  </div>
                  <div class="config-field-helper">How long to keep sampled device presence history. Used by the Presence page.</div>
                </td>
              </tr>
              <tr>
                <td class="config-field-label config-field-label-top">Connectivity event retention</td>
                <td class="config-field-value">
                  <div class="config-value-with-unit">
                    <input
                      name="connectivityRetention"
                      type="number"
                      min="1"
                      step="1"
                      class="form-control config-retention-input"
                      value={appConfig().ConnectivityRetention}
                      required
                    ></input>
                    <span class="config-field-unit">hours</span>
                  </div>
                  <div class="config-field-helper">How long to keep Online and Offline events.</div>
                  <div class="config-field-helper">Device-change events are retained while the device exists.</div>
                </td>
              </tr>
              <tr>
                <td></td>
                <td class="config-action-cell">
                  <button type="submit" class="btn btn-sm wyl-button">Save retention</button>
                  <span class={"config-save-status" + (error() ? " config-save-error" : "")} role="status">
                    {error() || status()}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </form>
      </div>
    </div>
  );
}

export default Retention;
