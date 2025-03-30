import { ALERT_DISPLAY_TIME } from '$lib/constants';
import { untrack } from 'svelte';

export type AlertType = 'error' | 'success';

export class Alert {
	message = $state('');
	type = $state<AlertType>('error');
	count = $state(0);
	id = $state(Math.random());
	#paused = false;
	#timeoutId: number;

	constructor(
		message: string,
		type: AlertType,
		private removeSelf: () => void,
		paused = false,
	) {
		this.message = message;
		this.type = type;
		this.count++;
		this.#paused = paused;
		this.#timeoutId = this.#removeTimer();
	}

	#removeTimer() {
		if (this.#paused) {
			return this.#timeoutId;
		}

		return setTimeout(() => {
			this.removeSelf();
		}, ALERT_DISPLAY_TIME);
	}

	duplicate() {
		clearTimeout(this.#timeoutId);

		this.count++;

		this.#timeoutId = this.#removeTimer();
	}

	isDuplicate(message: string, type: string): boolean {
		if (this.message === message && this.type === type) {
			return true;
		}

		return false;
	}

	remove() {
		clearTimeout(this.#timeoutId);
		this.removeSelf();
	}

	pause() {
		this.#paused = true;
		clearTimeout(this.#timeoutId);
	}

	unpause() {
		this.#paused = false;
		this.#timeoutId = this.#removeTimer();
	}
}

class Alerts {
	alerts = $state<Alert[]>([]);
	#paused = false;

	addAlert(message: string, type: AlertType) {
		for (const alert of this.alerts) {
			if (alert.isDuplicate(message, type)) {
				alert.duplicate();
				return;
			}
		}

		const removeAlert = () => {
			this.alerts = this.alerts.filter((alert) => alert.message !== message);
		};

		// use untrack to allow push to work in an effect
		untrack(() => {
			this.alerts.push(new Alert(message, type, removeAlert, this.#paused));
		});
	}

	addAlertFromFormError(form: unknown) {
		if (
			typeof form !== 'object' ||
			form === null ||
			!('message' in form) ||
			typeof form.message !== 'string'
		) {
			return null;
		}

		this.addAlert(form.message, 'error');
	}

	pauseAll() {
		this.#paused = true;
		for (const alert of this.alerts) {
			alert.pause();
		}
	}

	unpauseAll() {
		this.#paused = false;
		for (const alert of this.alerts) {
			alert.unpause();
		}
	}
}

export const alerts = new Alerts();
