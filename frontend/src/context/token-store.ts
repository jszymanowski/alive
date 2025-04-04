const TOKEN_NAME ='still-kicking-auth-token'

export const tokenStore = {
  subscribers: new Set<() => void>(),

  getToken: () => localStorage.getItem(TOKEN_NAME),

  setToken: (token: string | null) => {
    if (token) {
      localStorage.setItem(TOKEN_NAME, token);
    } else {
      localStorage.removeItem(TOKEN_NAME);
    }

    for (const callback of tokenStore.subscribers) {
      callback();
    }
  },

  subscribe: (callback: () => void) => {
    tokenStore.subscribers.add(callback);
    return () => tokenStore.subscribers.delete(callback);
  },
};
