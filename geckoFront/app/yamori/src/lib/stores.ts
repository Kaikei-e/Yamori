import { getContext, setContext } from 'svelte';
import { writable, type Writable } from 'svelte/store';

// store creation
export const session = getContextStore<App.Session>('__SESSION__');

// helper functions

export function setupContextStore<T>(key: string, data: T) {
  const store = setContext<Writable<T>>(key, writable(data));

  let initial = true;
  function update(data: T) {
    if (initial) {
      initial = false;
      return;
    }
    store.set(data);
  }
  return {
    store,
    update
  };
}

function getContextStore<T>(key: string): Writable<T> {
  const store: Writable<T> = {
    subscribe: (cb) => setup().subscribe(cb),
    set: (cb) => setup().set(cb),
    update: (cb) => setup().update(cb)
  };
  return store;

  function setup() {
    const { set, update, subscribe } = getContext<Writable<T>>(key);
    store.set = set;
    store.update = update;
    store.subscribe = subscribe;
    return store;
  }
}