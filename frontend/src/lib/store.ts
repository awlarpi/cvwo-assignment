import { create } from "zustand";
import { persist } from "zustand/middleware";

interface IStore {
  isLoggedIn: boolean;
  sessionId: string | null;
  userId: number | null;
  logIn: (sessionId: string, userId: number) => void;
  logOut: () => void;
  setSessionId: (sessionId: string | null) => void;
}

export const useStore = create(
  persist<IStore>(
    (set) => ({
      isLoggedIn: false,
      sessionId: null,
      userId: null,
      logIn: (sessionId: string, userId: number) =>
        set({ isLoggedIn: true, sessionId, userId }),
      logOut: () => set({ isLoggedIn: false, sessionId: null, userId: null }),
      setSessionId: (sessionId: string | null) => set({ sessionId }),
    }),
    {
      name: "login-storage", // name of the item in the storage (must be unique)
    }
  )
);
