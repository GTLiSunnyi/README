use std::net::UdpSocket;
use std::str;

fn main() {
    let listener = UdpSocket::bind("127.0.0.1:8080").expect("bind address failed");

    let mut buf = [0; 128];
    loop {
        let (size, src) = listener.recv_from(&mut buf).expect("recv data failed");
        println!(
            "size: {}, src: {}, data: {}",
            size,
            src,
            str::from_utf8(&buf[..size]).expect("convert byte failed")
        );

        listener
            .send_to(&buf[..size], src)
            .expect("send data failed");
    }
}
