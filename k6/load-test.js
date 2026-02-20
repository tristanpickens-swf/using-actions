import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = __ENV.TARGET_URL || 'http://localhost:8080';

export const options = {
    vus: 50,
    duration: '30s',
    thresholds: {
        http_req_duration: ['p(95)<500'],
        http_req_failed: ['rate<0.01'],
    },
};

export default function () {
    const url = `${BASE_URL}/contacts`;
    const payload = JSON.stringify({
        name: 'k6 User',
        phone: '555-k6-TEST',
    });
    const params = {
        headers: { 'Content-Type': 'application/json' },
    };
    const postRes = http.post(url, payload, params);
    check(postRes, {
        'is status 201': (r) => r.status === 201,
    });
    const getRes = http.get(url);
    check(getRes, { 'is status 200': (r) => r.status === 200 });
    sleep(1);
    }