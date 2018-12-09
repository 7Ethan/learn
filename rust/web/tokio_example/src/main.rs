#![deny(warnings)]

extern crate tokio;

use tokio::io;
use tokio::net::TcpListener;
use tokio::prelude::*;

use std::env;
use std::net::SocketAddr;

fn main() -> Result<(),Box<std::error::Error>>{
    let addr = env::args().nth(1).unwrap_or("127.0.0.1:8080".to_string());
    let addr = addr.parse::<SocketAddr>()?;

    let socket = TcpListener::bind(&addr)?;

    println!("Listening on {:?}", addr);

    let done = socket.incoming()
        .map_err(|e| eprintln!("failed to accept socket ,error :{:?}",e))
        .for_each(move |socket|{
            let (reader,writer) = socket.split();
            let amt = io::copy(reader,writer);

            let msg = amt.then(move |result|{
                match result{
                    Ok((amt,_,_)) => println!("wrote {} bytes.",amt),
                    Err(e) => eprintln!("Error : {}",e),
                }
                Ok(())
            });

            tokio::spawn(msg)
        });
    tokio::run(done);
    Ok(())
}

// fn main() {
//     // Bind the server's socket.
//     let addr = "127.0.0.1:8080".parse().unwrap();
//     let listener = TcpListener::bind(&addr).expect("unable to bind TCP listener");

//     // Pull out a stream of sockets for incoming connections
//     let server = listener
//         .incoming()
//         .map_err(|e| eprintln!("accept failed = {:?}", e))
//         .for_each(|sock| {
//             // Split up the reading and writing parts of the
//             // socket.
//             let (reader, writer) = sock.split();

//             // A future that echos the data and returns how
//             // many bytes were copied...
//             let bytes_copied = copy(reader, writer);

//             // ... after which we'll print what happened.
//             let handle_conn = bytes_copied
//                 .map(|amt| println!("wrote {:?} bytes", amt))
//                 .map_err(|err| eprintln!("IO error {:?}", err));

//             // Spawn the future as a concurrent task.
//             tokio::spawn(handle_conn)
//         });

//     // Start the Tokio runtime
//     tokio::run(server);
// }
