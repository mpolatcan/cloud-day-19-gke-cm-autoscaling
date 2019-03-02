import http from "k6/http";
import { sleep } from "k6";

export default function() {
    http.get("http://34.95.102.51/stress");
};