import axios from "axios";

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
