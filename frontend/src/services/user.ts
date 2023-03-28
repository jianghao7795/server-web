import { http } from "@/utils/request";

export function login(data: User.Login) {
  return http.post<User.UserInfo, User.Login>("/getImages", data);
}
