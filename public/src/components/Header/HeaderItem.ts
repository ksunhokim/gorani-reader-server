export enum SeenMode {
  EVERY = 1,
  LOGIN,
  LOGOUT,
}

export interface HeaderItem {
  name: string;
  callback?: any;
  endPoint?: string;
  seenMode: SeenMode;
}
