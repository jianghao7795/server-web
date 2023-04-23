import axios from "axios";

const service = axios.create();

export function Commits(page) {
  return service({
    url: `https://api.github.com/repos/JiangHaoCode/server-web/commits?page=${page}`,
    method: "get",
    headers: {
      "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
    },
  });
}

export function Members() {
  return service({
    url: "https://api.github.com/orgs/FLIPPED-AURORA/members",
    method: "get",
  });
}
