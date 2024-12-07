Features:

* [ ] create a PDF report

Marketer

* [x] Create new service

Expert

* [x] Create new parameter
* [x] Create new class
    * [ ] add a restriction for the class (to the IS)
* [ ] class approval for the new service by the expert
* [ ] In case of refusal of the proposed class - display a window with possible classes (IS)

Model

* [ ] Predicting a class for a new service
* [ ] Re-training the model when adding a new class/parameter

API:

* [x] POST /services
* [x] GET /services
* [x] POST /services/{id}/approve (if group is not in the body - it will approve group that was assigned earlier)
* [x] GET /services/{id}/proposed_groups
* [x] POST /parameters
* [x] GET /parameters
* [x] POST /groups
* [x] GET /groups
* [x] GET /report

Service {
  title: string,
  parameters: Parameter[],
  group: Group,
  created_at: timestamp,
  approved_at: timestamp,
  //approved_by: Expert,
}

Parameter {
  code: string,
  title: string,
} {
  new: boolean, 
}


Group {
  id: uint,
  title: string,
} {
  new: boolean,
  restrictions: string[], // -> OWL rules
}