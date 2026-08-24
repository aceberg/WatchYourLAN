import { createSignal, onMount, Show } from "solid-js"
import About from "../components/Config/About"
import Basic from "../components/Config/Basic"
import Influx from "../components/Config/Influx"
import Prometheus from "../components/Config/Prometheus"
import Retention from "../components/Config/Retention"
import Scan from "../components/Config/Scan"
import { refreshAppConfig } from "../functions/theme"

function Config() {
  const [loading, setLoading] = createSignal(true);
  const [error, setError] = createSignal("");

  onMount(() => {
    const scrollToHash = () => {
      const id = window.location.hash.slice(1);
      if (!id) {
        return;
      }

      document.getElementById(id)?.scrollIntoView();
    };

    requestAnimationFrame(scrollToHash);
    window.setTimeout(scrollToHash, 100);
    window.setTimeout(scrollToHash, 500);

    refreshAppConfig()
      .then(() => {
        setError("");
      })
      .catch(() => {
        setError("Settings could not be loaded. Check that the LANventory backend is reachable, then reload this page.");
      })
      .finally(() => {
        setLoading(false);
      });
  });

  return (
    <div class="settings-page">
      <header class="settings-page-header">
        <h1>Settings</h1>
        <p>Configure LANventory</p>
      </header>

      <Show when={!loading()} fallback={<div class="settings-status-panel" role="status">Loading settings</div>}>
        <Show when={!error()} fallback={<div class="settings-status-panel settings-status-error" role="alert">{error()}</div>}>
          <div class="settings-layout">
            <div class="settings-column">
              <section id="general" class="settings-section" aria-label="General">
                <Basic></Basic>
              </section>

              <section id="scanning" class="settings-section" aria-label="Scanning and database">
                <Scan></Scan>
              </section>

              <section id="data-retention" class="settings-section" aria-label="Data retention">
                <Retention></Retention>
              </section>
            </div>

            <div class="settings-column">
              <section id="integrations" class="settings-section" aria-label="Integrations">
                <div class="settings-section-heading">
                  <h2>Integrations</h2>
                  <p>Send LANventory metrics to external monitoring systems.</p>
                </div>
                <Influx></Influx>
                <Prometheus></Prometheus>
              </section>

              <section id="about" class="settings-section" aria-label="About">
                <About></About>
              </section>
            </div>
          </div>
        </Show>
      </Show>
    </div>
  )
}

export default Config
