import http from "k6/http";
import { sleep } from "k6";

export default function() {
    http.get("http://35.184.246.174");
};