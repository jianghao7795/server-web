import axios from "axios";
import serviceAxios from "@/utils/request";

const service = axios.create();

export function Commits(params) {
  return service({
    url: `https://api.github.com/repos/JiangHaoCode/server-web/commits`,
    method: "get",
    params,
  });
}

export function Members() {
  return service({
    url: "https://api.github.com/orgs/FLIPPED-AURORA/members",
    method: "get",
  });
}

export function getGithubCommitList(params) {
  return serviceAxios({
    url: "github/getGithubList",
    method: "get",
    params,
  });
}

export function createCommit() {
  return serviceAxios({
    url: "github/createGithub",
    method: "get",
  });
}
