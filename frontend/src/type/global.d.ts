declare namespace Global {
  export type Response<T> = {
    code: 7 | 0;
    data: T;
    msg: string;
  };
}
