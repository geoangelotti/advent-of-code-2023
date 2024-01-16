(ns day-04.part-1
  (:require [day-04.core :refer :all]))


(defn -main
  "Read the input and solve it"
  [& args]
  (println (day-04.core/process-part-1 (slurp "input.txt"))))
