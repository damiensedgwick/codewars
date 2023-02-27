fn main() {
   let bmi = bmi(170, 5.5);

    println!("bmi {:?}", bmi)
}

fn bmi(weight: u32, height: f32) -> &'static str {
    let bmi = weight as f32 / height.powi(2);

    match bmi {
        0.0..=18.5 => "Underweight",
        18.5..=25.0 => "Normal",
        25.0..=30.0 => "Overweight",
        _ => "Obese",
    }
}

#[cfg(test)]
mod tests {
    use crate::bmi;

    #[test]
    fn test_bmi() {
        assert_eq!(bmi(50, 1.80), "Underweight");
        assert_eq!(bmi(80, 1.80), "Normal");
        assert_eq!(bmi(90, 1.80), "Overweight");
        assert_eq!(bmi(110, 1.80), "Obese");
    }
}