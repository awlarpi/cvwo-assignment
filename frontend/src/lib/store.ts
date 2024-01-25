import { create } from "zustand";
import { persist } from "zustand/middleware";

interface IStore {
  isLoggedIn: boolean;
  sessionId: string | null;
  logIn: (sessionId: string) => void;
  logOut: () => void;
  setSessionId: (sessionId: string | null) => void;
}

export const useStore = create(
  persist<IStore>(
    (set) => ({
      isLoggedIn: false,
      sessionId: null,
      logIn: (sessionId: string) => set({ isLoggedIn: true, sessionId }),
      logOut: () => set({ isLoggedIn: false, sessionId: null }),
      setSessionId: (sessionId: string | null) => set({ sessionId }),
    }),
    {
      name: "login-storage", // name of the item in the storage (must be unique)
    }
  )
);
