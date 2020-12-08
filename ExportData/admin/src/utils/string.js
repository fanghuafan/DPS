export function confirmEnding(str, target) {
  var start = str.length - target.length;
  var arr = str.substr(start, target.length);
  if (arr == target) {
    return true;
  }
  return false;
}
