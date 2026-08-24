import { For, Show } from "solid-js";
import { A } from "@solidjs/router";

import {
  activityDescription,
  activityDeviceIcon,
  activityHostName,
  activityIcon,
  activityTone,
  relativeActivityTime,
} from "../functions/activity";
import type { HostEvent } from "../functions/exports";

type ActivityFeedProps = {
  events: HostEvent[];
  emptyText: string;
  hostExists?: (event: HostEvent) => boolean;
};

function ActivityFeed(props: ActivityFeedProps) {
  const canLinkHost = (event: HostEvent) => event.HostID > 0 && (props.hostExists?.(event) ?? false);

  return (
    <Show
      when={props.events.length > 0}
      fallback={<div class="activity-empty">{props.emptyText}</div>}
    >
      <div class="activity-feed" role="list">
        <For each={props.events}>{(event) =>
          <div class={"activity-row activity-row-" + activityTone(event.EventType)} role="listitem">
            <span class="activity-event-icon" aria-hidden="true">
              <i class={"bi " + activityIcon(event.EventType)}></i>
            </span>
            <span class="activity-host-icon" aria-hidden="true">
              <i class={"bi " + activityDeviceIcon(event)}></i>
            </span>
            <span class="activity-main">
              <Show
                when={canLinkHost(event)}
                fallback={<span class="activity-host-name">{activityHostName(event)}</span>}
              >
                <A href={"/host/" + event.HostID} class="activity-host-link">{activityHostName(event)}</A>
              </Show>
              <span class="activity-description">{activityDescription(event)}</span>
            </span>
            <time class="activity-time" dateTime={event.Date} title={event.Date}>
              {relativeActivityTime(event.Date)}
            </time>
          </div>
        }</For>
      </div>
    </Show>
  );
}

export default ActivityFeed;
