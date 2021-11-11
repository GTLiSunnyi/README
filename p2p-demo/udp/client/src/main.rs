use std::io::{self};
use std::net::UdpSocket;

fn main() {
    let client = UdpSocket::bind("127.0.0.1:8889").unwrap();

    client.connect("127.0.0.1:8080").unwrap();

    let mut buf = String::new();
    loop {
        let size = io::stdin().read_line(&mut buf).unwrap();
        println!("read: {}", size);
        let send_size = client.send(&buf.as_bytes()[..size]).unwrap();
        println!("size: {}, send: {}", size, send_size);
    }
}
