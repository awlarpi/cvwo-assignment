import { create } from "zustand";
import { persist } from "zustand/middleware";

interface IStore {
  isLoggedIn: boolean;
  logIn: () => void;
  logOut: () => void;
}

export const useStore = create(
  persist<IStore>(
    (set) => ({
      isLoggedIn: false,
      logIn: () => set({ isLoggedIn: true }),
      logOut: () => set({ isLoggedIn: false }),
    }),
    {
      name: "login-storage", // name of the item in the storage (must be unique)
    }
  )
);
