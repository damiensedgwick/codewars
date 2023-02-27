fn main() {
   let eighteen = how_many_lightsabers_do_you_own("Zach");
   let zero = how_many_lightsabers_do_you_own("Damien");

    println!("eighteen {:?}", eighteen);
    println!("zero: {:?}", zero);
}

fn how_many_lightsabers_do_you_own(name: &str) -> u8 {
    match name {
        "Zach" => 18,
        _ => 0
    }
}
