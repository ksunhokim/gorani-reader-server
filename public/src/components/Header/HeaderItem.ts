export enum SeenMode {
  EVERY = 1,
  LOGIN,
  LOGOUT,
}

export interface HeaderItem {
  name: string;
  endPoint: string;
  seenMode: SeenMode;
}
