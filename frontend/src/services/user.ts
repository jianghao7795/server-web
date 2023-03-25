import { http } from "@/utils/request";

export const aboutMe = () => {
    return http.get<API.ResponseAbout<API.AboutMe>>("/health");
  };