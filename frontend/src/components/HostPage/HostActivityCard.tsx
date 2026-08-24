import { createEffect, createSignal, onCleanup, onMount } from "solid-js";

import { apiGetHostActivity } from "../../functions/api";
import type { Host, HostEvent } from "../../functions/exports";
import ActivityFeed from "../ActivityFeed";

type HostActivityCardProps = {
  host: Host;
};

const hostActivityLimit = 10;

function HostActivityCard(props: HostActivityCardProps) {
  const [events, setEvents] = createSignal<HostEvent[]>([]);
  let requestId = 0;
  let refreshTimer = 0;

  const loadActivity = async (hostID: number) => {
    if (hostID < 1) {
      requestId++;
      setEvents([]);
      return;
    }

    const activeRequest = ++requestId;
    try {
      const nextEvents = await apiGetHostActivity(hostID, hostActivityLimit);
      if (activeRequest === requestId) {
        setEvents(nextEvents);
      }
    } catch {
      if (activeRequest === requestId) {
        setEvents([]);
      }
    }
  };

  createEffect(() => {
    loadActivity(props.host.ID);
  });

  onMount(() => {
    refreshTimer = window.setInterval(() => {
      loadActivity(props.host.ID);
    }, 60000);
  });

  onCleanup(() => {
    requestId++;
    window.clearInterval(refreshTimer);
  });

  const hostExists = (event: HostEvent) => event.HostID === props.host.ID && event.Mac === props.host.Mac;

  return (
    <section class="card wyl-panel activity-panel host-activity-panel" aria-labelledby="host-activity-title">
      <div class="card-header activity-panel-header">
        <div>
          <div id="host-activity-title" class="activity-panel-title">Recent events</div>
          <div class="activity-panel-subtitle">{props.host.Mac || "Waiting for host"}</div>
        </div>
      </div>
      <div class="card-body activity-panel-body">
        <ActivityFeed
          events={events()}
          emptyText="No recorded events yet"
          hostExists={hostExists}
        ></ActivityFeed>
      </div>
    </section>
  );
}

export default HostActivityCard;
