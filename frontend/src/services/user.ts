import { http } from "@/utils/request";

export function login(data: User.Login) {
  return http.post<Global.Response<User.CurrentUser>, User.Login>("/login", data);
}

export function getCurrentUser() {
  return http.get<Global.Response<User.CurrentUser>>("/getCurrentUser");
}

export function updateBackgroundImage(data: User.UpdateImages) {
  return http.put<Global.Response<string>>("/updateBackgroundImage", data);
}

export function registerUser(data: User.Register) {
  return http.post<Global.Response<User.CurrentUser>>("/register", data);
}
