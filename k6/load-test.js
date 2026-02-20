import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    vus: 50,
    duration: '30s',
    thresholds: {
        http_req_duration: ['p(95)<500'],
    },
};

export default function () {
    const url = 'http://app:8080/contacts';
    const payload = JSON.stringify({
        name: 'k6 User',
        phone: '555-k6-TEST',
    });
    const params = {
        headers: { 'Content-Type': 'application/json' },
    };
    const res = http.post(url, payload, params);
    check(res, {
        'is status 201': (r) => r.status === 201,
    });
    http.get(url);
    sleep(1);
    }